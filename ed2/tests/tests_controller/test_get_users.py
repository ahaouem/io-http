from controllers import CreateUserController


def test_get_users(filled_controller: CreateUserController) -> None:
    assert filled_controller.get_users() == filled_controller.get_users()
