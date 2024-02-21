from controllers import CreateUserController

def test_get_user_failure(filled_controller: CreateUserController, empty_controller: CreateUserController) -> None:
    assert empty_controller.get_user(0) == None
    assert filled_controller.get_user(5) == None
