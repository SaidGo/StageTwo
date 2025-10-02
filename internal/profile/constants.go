package profile

import "time"

// userValidationCodeExpire — TTL для кода валидации пользователя.
var userValidationCodeExpire = 10 * time.Minute
