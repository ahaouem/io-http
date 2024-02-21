from controllers import CreateUserController

def test_get_last_id(filled_controller: CreateUserController, empty_controller: CreateUserController) -> None:
    assert filled_controller.get_last_id() == 4
    assert empty_controller.get_last_id() == 0
