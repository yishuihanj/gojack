package config

type app struct {
	Name    string `toml:"name"`
	Desc    string `toml:"desc"`
	Version string `toml:"version"`
}

type logger struct {
	Level   string   `toml:"level"`   //日志级别：error > info > trace
	Console bool     `toml:"console"` //是否打印日志到控制台
	Logfile *logFile `toml:"file"`    //日志文件属性
}
type logFile struct {
	Path      string `toml:"path"`       //日志输出路径，默认当前执行文件下的logs目录下
	Rotation  int    `toml:"rotation"`   //文件切割周期（天）
	Monitor   int    `toml:"monitor"`    //文件监控周期（天）
	Retain    int    `toml:"retain"`     //日志保留周期（天）
	LimitSize int    `toml:"limit_size"` //单个文件大小上限（M）
}
