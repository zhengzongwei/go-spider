/*
 * Copyright (c) 2023-2024 zhengzongwei<zhengzongwei@foxmail.com>
 *
 * Licensed under the Mulan Permissive Software License，Version 2;
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 		http://license.coscl.org.cn/MulanPSL2
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package utils

import (
	"fmt"
	"time"
)

func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimeStamp() int64 {
	currentTimeStamp := time.Now().Unix()
	return currentTimeStamp
}

func TimeStamp2TimeStr(timestamp int64) string {

	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

func TimeStr2TimeStamp(timeStr string) int64 {
	// Parse the time string
	Time, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		// Handle parsing error, for simplicity, returning 0
		fmt.Println("Error parsing time:", err)
		return 0
	}

	return Time.Unix()
}
