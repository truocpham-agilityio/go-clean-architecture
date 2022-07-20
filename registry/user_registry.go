package registry

import (
	"go-clean-architecture/interface/controller"
	ip "go-clean-architecture/interface/presenter"
	ir "go-clean-architecture/interface/repository"
	"go-clean-architecture/usecase/interactor"
	up "go-clean-architecture/usecase/presenter"
	ur "go-clean-architecture/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
