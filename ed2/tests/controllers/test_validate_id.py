from controllers import CreateUserController

def test_validate_id(empty_controller: CreateUserController) -> None:
    assert empty_controller.validate_id(1) == False
