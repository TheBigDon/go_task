package model_test

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name 	 string
		u func() *model.User
		isValid  bool
	} {
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty mobile",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Mobile = ""

				return u
			},
			isValid: false,
		},
		{
			name: "empty name",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = ""

				return u
			},
			isValid: false
		},
		{
			name: "invalid mobile",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Mobile = "invalid"

				return u
			},
			isValid: false,
		},
		{
			name: "invalid name",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Name = "invalid"

				return u
			},
			isValid: false,
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}