module github.com/sulicat/drawsaface

go 1.23.3

require (
	github.com/abema/go-mp4 v1.2.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/sulicat/goboi v0.0.0-20241210021630-930594058b69
	golang.org/x/term v0.27.0
)

require (
	github.com/google/uuid v1.1.2 // indirect
	golang.org/x/sys v0.28.0 // indirect
)

replace github.com/sulicat/goboi v0.0.0-20241210021630-930594058b69 => ../../git/goboi
