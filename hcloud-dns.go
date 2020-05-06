/*

	Created by Vladimir Smagin, 2020
	http://blindage.org 21h@blindage.org

	Project page: https://git.blindage.org/21h/hcloud-dns

	License GPL v3

*/

package hclouddns

// New instance
func New(t string) *HCloudDNS {
	return &HCloudDNS{
		token: t,
	}
}
