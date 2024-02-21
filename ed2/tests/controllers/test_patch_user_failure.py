from controllers import CreateUserController

def test_patch_user_failure(filled_controller: CreateUserController) -> None:
    user = {
        "firstName": "John",
        "lastName": "Doe",
        "birth_year": 1990,
        "group": "user"
    }
    assert filled_controller.patch_user(5, user) == False
    assert filled_controller.patch_user(0, {}) == False
    assert filled_controller.patch_user(0, {"firstName": "John"}) == False
