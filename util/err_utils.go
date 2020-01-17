package util

import "github.com/rs/zerolog/log"

/**
处理错误的工具类
*/

// 处理数据库错误
func handleDbError(err error) error {
	if err != nil {
		log.Error().Msg("DB ERR: " + err.Error())
		return err
	}
	return nil
}