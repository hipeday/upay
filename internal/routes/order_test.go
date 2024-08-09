package routes

import "testing"

func Test_getContractTransaction(t *testing.T) {
	type args struct {
		contractAddress string
		accountAddress  string
	}
	var tests []struct {
		name string
		args args
	}

	test := struct {
		name string
		args args
	}{
		name: "获取支付记录",
		args: args{
			contractAddress: "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj",
			accountAddress:  "TYYi4mhbkUwcezbnXP5UyKUbcS15Kx1qt8",
		},
	}
	tests = append(tests, test)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//trongrid.GetContractTransaction(tt.args.contractAddress, tt.args.accountAddress)
		})
	}
}
