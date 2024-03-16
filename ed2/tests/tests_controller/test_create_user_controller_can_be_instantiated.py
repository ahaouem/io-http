from controllers import CreateUserController


def test_create_user_controller_can_be_instantiated(empty_controller: CreateUserController) -> None:
    assert isinstance(empty_controller, CreateUserController)
