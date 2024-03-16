from controllers import CreateUserController


def test_remove_user(filled_controller: CreateUserController) -> None:
    assert filled_controller.remove_user(0) == True
    assert filled_controller.get_user(0) == None
