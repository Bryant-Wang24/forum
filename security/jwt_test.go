package security

import (
	"testing"

	"example.com/gin_forum/utils"
)

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT("jack", "jack@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("token: %v\n", token)
}

func TestVerifyJWT(t *testing.T) {
	claim, valid, err := VerifyJWT("eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTA1NjcwMTMsImlhdCI6MTcxMDQ4MDYxMywidXNlciI6eyJlbWFpbCI6ImphY2tAZ21haWwuY29tIiwidXNlcm5hbWUiOiJqYWNrIn19.CHRJXrKWVIQ5aSpj3RMQLTAMaDNoRR9afEqEvbB88s5ePMlNg55AWpafAD57K6gMPukiOF3e2vrU64v8DPHjeVvn10-CuidliQgsZ_fGoCe4skbesASlYJH-gE7BlQMJMNFA2a8y3gVgPoew_cRD7Un84Qm0AmW8v8Afx47bZw7hAfssokeQIpoNAyp2WfqkfHknhV7nLqXtiHnCPy7rQeCluM4XnooTE6YvQrDukr0oRTTBSBb71rIpstGygBaLdrEh6pr9a1YJgODKzkfnauDl_pENMM5fwMfAl1b9ABXOhwNDb-XRt1WkBPrGEdcTaW8kznzgQ4GWdSSh1LkPog")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("verify jwt: %v\n,claim: %v\n", valid, utils.JsonMarshal(claim))
}
