package main

type CertificateCmd struct {
	Parse *ParseCmd `arg:"subcommand:parse"`
}

func doCertificateCmd(cmd *CertificateCmd) {
	if cmd == nil {
		logger.Fatalf("cmd cannot be nil")
	}

	if cmd.Parse != nil {
		doParseCertificate(cmd.Parse)
		return
	}

	logger.Errorf(specifySubcommand)
}
