// Copyright 2012 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Distro: Debian

package osutil

import "github.com/tredoe/osutil/executil"

// 'apt' is for the terminal and gives beautiful output.
// 'apt-get' and 'apt-cache' are for scripts and give stable, parsable output.

const pathDeb = "/usr/bin/apt-get"

// ManagerDeb is the interface to handle the package manager of Linux systems based at Debian.
type ManagerDeb struct{}

func (p ManagerDeb) Install(name ...string) error {
	args := []string{pathDeb, "install", "-y"}

	return executil.RunToStd(nil, sudo, append(args, name...)...)
}

func (p ManagerDeb) Remove(name ...string) error {
	args := []string{pathDeb, "remove", "-y"}

	return executil.RunToStd(nil, sudo, append(args, name...)...)
}

func (p ManagerDeb) Purge(name ...string) error {
	args := []string{pathDeb, "purge", "-y"}

	return executil.RunToStd(nil, sudo, append(args, name...)...)
}

func (p ManagerDeb) Update() error {
	return executil.RunToStd(nil, sudo, pathDeb, "update", "-qq")
}

func (p ManagerDeb) Upgrade() error {
	return executil.RunToStd(nil, sudo, pathDeb, "upgrade", "-y")
}

func (p ManagerDeb) Clean() error {
	err := executil.RunToStd(nil, sudo, pathDeb, "autoremove", "-y")
	if err != nil {
		return err
	}

	return executil.RunToStd(nil, sudo, pathDeb, "clean")
}