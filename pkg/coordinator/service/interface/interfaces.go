package interfaces

import (
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

type CoordinatorSVCInter interface {
	SignupSVC(p *cpb.Signup) (*cpb.SignupResponce,error)
	VerifySVC(p *cpb.Verify) (*cpb.VerifyResponce, error)
	UserLogin(p *cpb.CoorinatorLogin)(*cpb.CordinatorLoginResponce,error)
	AddPackageSVC(p *cpb.AddPackage)(*cpb.AddPackageResponce,error)
	AddDestinationSVC(p *cpb.AddDestination)(*cpb.AddDestinationResponce,error)
	AddActivitySVC(p *cpb.AddActivity)(*cpb.AddActivityResponce,error)
}
