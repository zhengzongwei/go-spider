/*
 * Copyright (c) 2023 zhengzongwei
 *
 * Licensed under the Mulan Permissive Software License，Version 2;
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://license.coscl.org.cn/MulanPSL2
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

package utils

import (
	"fmt"
	"time"
)

//const TimeFormat = "2006-01-02 15:04:05"

func NowTime() {
	nowTime := time.Now()
	fmt.Printf("current time:%v\n", nowTime)
	return
}
