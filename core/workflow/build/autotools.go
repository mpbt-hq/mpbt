// SPDX-License-Identifier: AGPL-3.0-or-later
package build

import (
	"fmt"
)

type AutotoolsBuilder struct {
	BuilderBase
}

func (ab *AutotoolsBuilder) RunPrepare() error {
	return ab.ExecInSourceDir([]string{"./autogen.sh"})
}

func (ab *AutotoolsBuilder) RunConfigure() error {
	args := []string{"./configure",
		fmt.Sprintf("--prefix=%s", ab.Package.GetInstallPrefix())}

	autoconf_args := ab.Package.GetStrList("autoconf-args")
	autoconf_extra_args := ab.Package.GetStrList("autoconf-extra-args")

	args = append(args, autoconf_args...)
	args = append(args, autoconf_extra_args...)

	return ab.ExecInSourceDir(args)
}

func (ab *AutotoolsBuilder) RunBuild() error {
	return ab.ExecInSourceDir([]string{"make", fmt.Sprintf("-j%d", ab.Package.GetParallel())})
}

func (ab *AutotoolsBuilder) RunInstall() error {
	return ab.ExecInSourceDir([]string{"make", "install", "DESTDIR=" + ab.Package.GetDestdir()})
}

func (ab *AutotoolsBuilder) RunClean() error {
	return ab.ExecInSourceDir([]string{"make", "clean"})
}
