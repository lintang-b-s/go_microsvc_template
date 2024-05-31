package pkg

import (
	"fmt"
	"lintang/go_hertz_template/config"
	"os"
	"regexp"

	"github.com/cloudwego/hertz/pkg/app/server/binding"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func PasswordValidator() *binding.ValidateConfig {
	// golang gabisa pake regex '^?=.*[0-9]?=.*[a-z]?=.*[A-Z]?=.*\\W(?!.* )*$'
	passwordVDConfig := &binding.ValidateConfig{}
	passwordVDConfig.MustRegValidateFunc("password", func(args ...interface{}) error {
		password, _ := args[0].(string)
		lengthRegex := regexp.MustCompile(`.{8,}`)        // Memeriksa panjang minimal 8 karakter
		upperRegex := regexp.MustCompile(`[A-Z]`)         // Memeriksa adanya huruf besar
		lowerRegex := regexp.MustCompile(`[a-z]`)         // Memeriksa adanya huruf kecil
		numberRegex := regexp.MustCompile(`\d`)           // Memeriksa adanya angka
		specialCharRegex := regexp.MustCompile(`[@#$%*]`) // Memeriksa adanya karakter spesial

		isMatch := lengthRegex.MatchString(password) &&
			upperRegex.MatchString(password) &&
			lowerRegex.MatchString(password) &&
			numberRegex.MatchString(password) &&
			specialCharRegex.MatchString(password)
		if !isMatch {
			return fmt.Errorf("password harus terdiri dari minimal 8 karakter, 1 uppercase, 1 lowercase, 1 digit, 1 karakter spesial")
		}
		return nil
	})
	return passwordVDConfig
}

var lg *zap.Logger

// pake hertzlogger gak kayak pake uber/zap logger beneran
func InitZapLogger(cfg *config.Config) *hertzzap.Logger {
	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	productionCfg.EncodeDuration = zapcore.SecondsDurationEncoder
	productionCfg.EncodeCaller = zapcore.ShortCallerEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// log encooder (json for prod, console for dev)
	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)
	// loglevel
	logDevLevel := zap.NewAtomicLevelAt(zap.DebugLevel)
	logLevelProd := zap.NewAtomicLevelAt(zap.InfoLevel)

	//write sycer
	writeSyncerStdout, writeSyncerFile := getLogWriter(cfg.MaxBackups, cfg.MaxAge)

	prodCfg := hertzzap.CoreConfig{
		Enc: fileEncoder,
		Ws:  writeSyncerFile,
		Lvl: logLevelProd,
	}

	devCfg := hertzzap.CoreConfig{
		Enc: consoleEncoder,
		Ws:  writeSyncerStdout,
		Lvl: logDevLevel,
	}
	logsCores := []hertzzap.CoreConfig{
		prodCfg,
		devCfg,
	}
	coreConsole := zapcore.NewCore(consoleEncoder, writeSyncerStdout, logDevLevel)
	coreFile := zapcore.NewCore(fileEncoder, writeSyncerFile, logLevelProd)
	core := zapcore.NewTee(
		coreConsole,
		coreFile,
	)
	lg = zap.New(core)
	zap.ReplaceGlobals(lg)

	prodAndDevLogger := hertzzap.NewLogger(hertzzap.WithZapOptions(zap.WithFatalHook(zapcore.WriteThenPanic)),
		hertzzap.WithCores(logsCores...))

	return prodAndDevLogger
}

func getLogWriter(maxBackup, maxAge int) (writeSyncerStdout zapcore.WriteSyncer, writeSyncerFile zapcore.WriteSyncer) {
	file := zapcore.AddSync(&lumberjack.Logger{
		Filename: "./logs/app.log",

		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	})
	stdout := zapcore.AddSync(os.Stdout)

	return stdout, file
}

type ValidateError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *ValidateError) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return e.ErrType + ": expr_path=" + e.FailField + ", cause=invalid"
}

type BindError struct {
	ErrType, FailField, Msg string
}

// Error implements error interface.
func (e *BindError) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return e.ErrType + ": expr_path=" + e.FailField + ", cause=invalid"
}

func CreateCustomValidationError() *binding.ValidateConfig {
	validateConfig := &binding.ValidateConfig{}
	validateConfig.SetValidatorErrorFactory(func(failField, msg string) error {
		err := ValidateError{
			ErrType:   "validateErr",
			FailField: "[validateFailField]: " + failField,
			Msg:       msg,
		}

		return &err
	})
	return validateConfig
}
