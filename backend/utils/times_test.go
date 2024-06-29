/*
 * Copyright (c) 2024 zhengzongwei<zhengzongwei@foxmail.com>
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
	"testing"
)

func TestTimeStr2TimeStamp(t *testing.T) {
	TimeStamp := TimeStr2TimeStamp("2024-01-03 16:57:17")
	fmt.Printf("%d", TimeStamp)
}
func TestTimeStamp2TimeStr(t *testing.T) {
	TimeStr := TimeStamp2TimeStr(1704272237)
	fmt.Printf("%s", TimeStr)
}
func TestTimeStamp(t *testing.T) {
	time := TimeStamp()
	fmt.Printf("%d", time)
}
func TestNowTime(t *testing.T) {
	NowTime()
}
