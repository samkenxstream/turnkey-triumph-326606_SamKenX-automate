package main

import (
	"os"
	"testing"

	"github.com/chef/automate/lib/io/fileutils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const (
	RemoteFilePath = "198.51.100.0:/home/ec2-user/certs/public.pem"
	LocalFilePath  = "/home/ec2-user/certs/public.pem"
	ValidIP        = "198.51.100.0"
	FileContent    = `-----BEGIN CERTIFICATE-----
MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
U3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw
NjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE
ChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp
ZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3
DQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf
8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN
+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0
X9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa
K4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA
1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G
A1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR
zt0fhvRbVazc1xDCDqmI56FspGowaaBCDEfGA1UEBhMCVVMxJTAjBgNVBAoTHFN0
YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD
bGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w
DQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3
L7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D
eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8xY5Z=
-----END CERTIFICATE-----`
	ValidCertPath = "./certRotate.go"
)

func TestIsRemotePath(t *testing.T) {
	c := certRotateFlow{FileUtils: mockFS()}

	type testCaseInfo struct {
		testCaseDescription string
		input               string
		expected            bool
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "Valid remote path",
			input:               RemoteFilePath,
			expected:            true,
		},
		{
			testCaseDescription: "Local path instead of remote path",
			input:               LocalFilePath,
			expected:            false,
		},
		{
			testCaseDescription: "Invalid remote path 1",
			input:               "/home/ec2-user/certs/public.pem198.51.100.0",
			expected:            false,
		},
		{
			testCaseDescription: "Invalid remote path 2",
			input:               "198.51.100.0/home/ec2-user/certs/public.pem",
			expected:            false,
		},
		{
			testCaseDescription: "Invalid remote path 3",
			input:               "\n   198.51.100.0:/home/ec2-user/certs/public.pem",
			expected:            false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {
			actual := c.IsRemotePath(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestGetIPV4(t *testing.T) {
	c := certRotateFlow{FileUtils: mockFS()}

	type testCaseInfo struct {
		testCaseDescription string
		input               string
		expected            string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "Valid IP V4",
			input:               RemoteFilePath,
			expected:            ValidIP,
		},
		{
			testCaseDescription: "Valid IP V4 but invalid remote path 1",
			input:               "/home/ec2-user/certs/public.pem:127.0.0.1",
			expected:            "127.0.0.1",
		},
		{
			testCaseDescription: "Valid IP V4 but invalid remote path 2",
			input:               "/home/ec2-user/:0.0.0.0:/certs/public.pem",
			expected:            "0.0.0.0",
		},
		{
			testCaseDescription: "Invalid IP v4 and valid path",
			input:               "256.256.256.256:/home/ec2-user/certs/public.pem",
			expected:            "",
		},
		{
			testCaseDescription: "Invalid IP v4 and invalid path",
			input:               "/home/ec2-user/certs/public.pem:1.2.3",
			expected:            "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {
			actual := c.GetIPV4(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestGetRemoteFileDetails(t *testing.T) {
	c := certRotateFlow{FileUtils: mockFS()}

	type testCaseInfo struct {
		testCaseDescription    string
		input                  string
		isError                bool
		expectedErrorMessage   string
		expectedRemoteFilePath string
		expectedFileName       string
		expectedHostIP         string
		clusterIPs             []string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription:    "Valid remote path",
			input:                  RemoteFilePath,
			isError:                false,
			expectedErrorMessage:   "",
			expectedRemoteFilePath: LocalFilePath,
			expectedFileName:       "public.pem",
			expectedHostIP:         ValidIP,
			clusterIPs:             []string{ValidIP},
		},
		{
			testCaseDescription:    "Invalid Remote Path - Local path",
			input:                  LocalFilePath,
			isError:                true,
			expectedErrorMessage:   " is not a valid IPv4 address",
			expectedRemoteFilePath: "",
			expectedFileName:       "",
			expectedHostIP:         "",
			clusterIPs:             []string{ValidIP},
		},
		{
			testCaseDescription:    "Invalid Remote Path - Colon missing",
			input:                  "198.51.100.0/home/ec2-user/certs/public.pem",
			isError:                true,
			expectedErrorMessage:   "Invalid remote path: 198.51.100.0/home/ec2-user/certs/public.pem",
			expectedRemoteFilePath: "",
			expectedFileName:       "",
			expectedHostIP:         "",
			clusterIPs:             []string{ValidIP},
		},
		{
			testCaseDescription:    "Invalid Remote Path - No filename",
			input:                  "198.51.100.0:/home/ec2-user/certs/public/",
			isError:                false,
			expectedErrorMessage:   "",
			expectedRemoteFilePath: "/home/ec2-user/certs/public",
			expectedFileName:       "public",
			expectedHostIP:         ValidIP,
			clusterIPs:             []string{ValidIP},
		},
		{
			testCaseDescription:    "Invalid Remote Path - Reverse",
			input:                  "/home/ec2-user/certs/public/:198.51.100.0",
			isError:                false,
			expectedErrorMessage:   "",
			expectedRemoteFilePath: ValidIP,
			expectedFileName:       ValidIP,
			expectedHostIP:         ValidIP,
			clusterIPs:             []string{ValidIP},
		},
		{
			testCaseDescription:    "Invalid Remote Path - Empty Path",
			input:                  ValidIP + ":",
			isError:                true,
			expectedErrorMessage:   "Invalid remote path: " + ValidIP + ":",
			expectedRemoteFilePath: "",
			expectedFileName:       "",
			expectedHostIP:         "",
			clusterIPs:             []string{ValidIP},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {
			infra := &AutomteHAInfraDetails{}
			infra.Outputs.AutomatePrivateIps.Value = tc.clusterIPs
			remoteFilePathRes, fileNameRes, hostIPRes, err := c.GetRemoteFileDetails(tc.input, infra)
			if tc.isError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErrorMessage, err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedRemoteFilePath, remoteFilePathRes)
			assert.Equal(t, tc.expectedFileName, fileNameRes)
			assert.Equal(t, tc.expectedHostIP, hostIPRes)
		})
	}
}

func TestGetCerts(t *testing.T) {
	c := certRotateFlow{FileUtils: mockFS()}

	infra := &AutomteHAInfraDetails{}

	type testCaseInfo struct {
		testCaseDescription string
		flagsObj            flags
		rootCaWant          string
		publicCertWant      string
		privateCertWant     string
		adminCertWant       string
		adminKeyWant        string
		isError             bool
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "All paths given and flag is automate service",
			flagsObj: flags{
				automate:        true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				rootCAPath:      ValidCertPath,
			},
			rootCaWant:      FileContent,
			publicCertWant:  FileContent,
			privateCertWant: FileContent,
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         false,
		},
		{
			testCaseDescription: "All paths given except root-ca flag is automate service and node flag given",
			flagsObj: flags{
				automate:        true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				node:            "ip-given",
			},
			rootCaWant:      "",
			publicCertWant:  FileContent,
			privateCertWant: FileContent,
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         false,
		},
		{
			testCaseDescription: "All paths given and flag is opensearch service",
			flagsObj: flags{
				opensearch:      true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				rootCAPath:      ValidCertPath,
				adminCertPath:   ValidCertPath,
				adminKeyPath:    ValidCertPath,
			},
			rootCaWant:      FileContent,
			publicCertWant:  FileContent,
			privateCertWant: FileContent,
			adminCertWant:   FileContent,
			adminKeyWant:    FileContent,
			isError:         false,
		},
		{
			testCaseDescription: "All paths empty and flag for automate service",
			flagsObj:            flags{},
			rootCaWant:          "",
			publicCertWant:      "",
			privateCertWant:     "",
			adminCertWant:       "",
			adminKeyWant:        "",
			isError:             true,
		},
		{
			testCaseDescription: "some invalid paths given and flag is automate service",
			flagsObj: flags{
				automate:        true,
				privateCertPath: ValidCertPath,
				publicCertPath:  "./xyx-cert.go",
				rootCAPath:      ValidCertPath,
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
		{
			testCaseDescription: "All paths given but invalid (file not exist in (f.s)and flag is automate service",
			flagsObj: flags{
				automate:        true,
				privateCertPath: "./xyz.go",
				publicCertPath:  "./xyz.go",
				rootCAPath:      "./xyx.go",
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
		{
			testCaseDescription: "All paths given except root-ca and flag is automate service",
			flagsObj: flags{
				automate:        true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
		{
			testCaseDescription: "All paths given but root-ca path is invalid(file not exist) flag is automate service",
			flagsObj: flags{
				automate:        true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				rootCAPath:      "./xyx-cert.go",
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
		{
			testCaseDescription: "Some mandatory path not given and flag is opensearch service",
			flagsObj: flags{
				opensearch:      true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				rootCAPath:      ValidCertPath,
				adminCertPath:   ValidCertPath,
				adminKeyPath:    "",
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
		{
			testCaseDescription: "Invalid adminCert path and flag is opensearch service",
			flagsObj: flags{
				opensearch:      true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				rootCAPath:      ValidCertPath,
				adminCertPath:   "./xyz-cert.go",
				adminKeyPath:    ValidCertPath,
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
		{
			testCaseDescription: "Invalid adminKey path and flag is opensearch service",
			flagsObj: flags{
				opensearch:      true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				rootCAPath:      ValidCertPath,
				adminCertPath:   ValidCertPath,
				adminKeyPath:    "./xyz-cert.go",
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {
			certsGot, err := c.getCerts(infra, &tc.flagsObj)
			if tc.isError {
				assert.Error(t, err)
				assert.Nil(t, certsGot)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, certsGot)
			}
			if certsGot != nil {
				assert.Equal(t, tc.rootCaWant, certsGot.rootCA)
				assert.Equal(t, tc.publicCertWant, certsGot.publicCert)
				assert.Equal(t, tc.privateCertWant, certsGot.privateCert)
				assert.Equal(t, tc.adminCertWant, certsGot.adminCert)
				assert.Equal(t, tc.adminKeyWant, certsGot.adminKey)
			}
		})
	}
}

func mockFS() *fileutils.MockFileSystemUtils {
	return &fileutils.MockFileSystemUtils{
		ReadFileFunc: func(filename string) ([]byte, error) {
			if _, err := os.Stat(filename); err == nil {
				// path/to/whatever exists
				return []byte(FileContent), nil
			} else if errors.Is(err, os.ErrNotExist) {
				// path/to/whatever does *not* exist
				return []byte{}, err
			} else {
				return []byte{}, err
			}
		},
	}
}