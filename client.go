package main

type ClientCmd struct {
	GetCertificate *GetCertificateCmd `arg:"subcommand:get-certificate"`
}

func doClientCmd() {
	if args.Client.GetCertificate != nil {
		doGetCertificateCmd()
	} else {
		logger.Fatalf(specifySubcommand)
	}
}
