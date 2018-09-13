/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package driver

import (
	"os"

	"golang.org/x/sys/unix"
)

// Lchmod changes the mode of a file not following symlinks.
func (d *driver) Lchmod(path string, mode os.FileMode) error {
	// On Linux, file mode is not supported for symlinks,
	// and fchmodat() does not support AT_SYMLINK_NOFOLLOW,
	// so symlinks need to be skipped entirely.
	if st, err := os.Stat(path); err == nil && st.Mode()&os.ModeSymlink != 0 {
		return nil
	}

	return unix.Fchmodat(unix.AT_FDCWD, path, uint32(mode), 0)
}
