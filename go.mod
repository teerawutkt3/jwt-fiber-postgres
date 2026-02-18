module fiber-poc-api

go 1.22

require (
	github.com/gofiber/fiber/v2 v2.52.6 // Patch update: แก้บั๊กเล็กน้อย
	github.com/golang-jwt/jwt/v5 v5.2.2 // Security: อัปเดตล่าสุด
	github.com/google/uuid v1.6.0 // Performance: ดีขึ้นเล็กน้อย
	github.com/jackc/pgx/v4 v4.18.3 // หมายเหตุ: GORM ใช้ pgx/v5 ถ้าคุณไม่ได้ใช้ v4 โดยตรงในโค้ด ลบอันนี้ออกได้เลย
	github.com/spf13/viper v1.19.0 // Security & Features: รองรับ config format ใหม่ๆ ดีขึ้น
	golang.org/x/crypto v0.32.0 // Critical Security: สำคัญที่สุดต้องเป็นเวอร์ชันใหม่เสมอ
	gorm.io/driver/postgres v1.5.11 // Update: ให้เข้ากับ pgx/v5 ล่าสุด
	gorm.io/gorm v1.25.12 // Update: Bug fixes
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gofiber/contrib/jwt v1.1.2
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.55.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require github.com/MicahParks/keyfunc/v2 v2.1.0 // indirect
