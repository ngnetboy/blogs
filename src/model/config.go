package model

import (
	"flag"
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var Conf *Configuration

type Configuration struct {
	LogLevel     string
	LogFile      string
	LogMaxSize   int
	LogMaxBackUp int
	LogMaxAge    int
	Addr         string
	DB           string
	PageNum      int
}

var (
	flagLogLevel     = flag.String("log.level", "info", "debug|info|warning|error")
	flagLogFile      = flag.String("log.file", "", "log file, write to console if it is empty")
	flagLogMaxSize   = flag.Int("log.maxsize", 10, "log file max size in megabytes")
	flagLogMaxBackup = flag.Int("log.maxbackup", 3, "log file max backup count")
	flagLogMaxAge    = flag.Int("log.maxage", 30, "log file max age in days")
	flagAddr         = flag.String("addr", "0.0.0.0:6688", "address to listen, ip:port")
	// flagDB           = flag.String("db", "host=127.0.0.1 user=postgres dbname=uac sslmode=disable", "database connect string")
	flagDB = flag.String("db", "blogs.db", "database connect string")
)

var Models = []interface{}{
	&Category{}, &User{}, &Article{}, &Tag{},
}

func initLog() {
	if level, err := log.ParseLevel(Conf.LogLevel); err == nil {
		log.SetLevel(level)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if len(Conf.LogFile) == 0 {
		log.SetOutput(os.Stdout)
		return
	}

	dir := path.Dir(Conf.LogFile)
	os.Mkdir(dir, 0755)

	log.SetOutput(&lumberjack.Logger{
		Filename:   Conf.LogFile,
		MaxSize:    Conf.LogMaxSize,
		MaxBackups: Conf.LogMaxBackUp,
		MaxAge:     Conf.LogMaxAge,
	})
}

func LoadConf() {
	flag.Parse()

	Conf = &Configuration{
		LogLevel:     *flagLogLevel,
		LogFile:      *flagLogFile,
		LogMaxSize:   *flagLogMaxSize,
		LogMaxBackUp: *flagLogMaxBackup,
		LogMaxAge:    *flagLogMaxAge,
		Addr:         *flagAddr,
		DB:           *flagDB,
		PageNum:      10,
	}

	if *flagLogMaxAge < 0 {
		Conf.LogMaxAge = 10
	}
	if *flagLogMaxBackup < 0 {
		Conf.LogMaxBackUp = 3
	}
	if *flagLogMaxSize < 0 {
		Conf.LogMaxSize = 10
	}

	initLog()

}
