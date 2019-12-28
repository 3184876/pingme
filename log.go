package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"time"

	"github.com/noobly314/pingme/httping"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init_log() {
	formatter := &logrus.TextFormatter{
		DisableTimestamp: true,
	}
	log.SetFormatter(formatter)
	log.Out = os.Stdout
}

func logHttping(stats httping.Stats, err error, address string) {
	if err == nil {
		fmt.Println("scheme: ", stats.Scheme)
		if stats.Scheme == "http" {
			log.Info("    Host               ", parseInput(address))
			log.Info("    DNS Lookup         ", stats.DNS/1e6)
			log.Info("    TCP                ", stats.TCP/1e6)
			log.Info("    Total              ", stats.Total/1e6)
		} else if stats.Scheme == "https" {
			log.Info("    Host               ", parseInput(address))
			log.Info("    DNS Lookup         ", stats.DNS/1e6)
			log.Info("    TCP                ", stats.TCP/1e6)
			log.Info("    TLS                ", stats.TLS/1e6)
			log.Info("    Process            ", stats.Process/1e6)
			log.Info("    Transfer           ", stats.Transfer/1e6)
			log.Info("    Total              ", stats.Total/1e6)
		}
	}
}

func logTcping(code int, address string) {
	if code == 0 {
		log.Info("    TCP     OPEN      ", address)
	} else if code == 1 {
		log.Warn("    TCP     CLOSED    ", address)
	} else if code == 2 {
		log.Warn("    TCP     ERROR     ", address)
	}
}

func logPing(dst *net.IPAddr, dur time.Duration, err error) {
	if err != nil {
		match, _ := regexp.MatchString("operation not permitted", err.Error())
		if match {
			log.Warn(fmt.Sprintf("    ICMP    ERROR     No privileges"))
		} else {
			log.Warn(fmt.Sprintf("    ICMP    ERROR     %s", dst.String()))
		}
		return
	}
	log.Info(fmt.Sprintf("    ICMP    OPEN      %s    %s ms", dst.String(), fmt.Sprintf("%.1f", float64(dur.Microseconds())/1000)))
}

func logMtr(hops []string, address string) {
	for _, h := range hops {
		log.Info("    MTR     ", h)
	}
}
