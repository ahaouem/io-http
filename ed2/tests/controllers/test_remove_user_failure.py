from controllers import CreateUserController

def test_remove_user_failure(filled_controller: CreateUserController) -> None:
    assert filled_controller.remove_user(5) == False
    assert filled_controller.remove_user(0) == True
