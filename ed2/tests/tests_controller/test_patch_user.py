from controllers import CreateUserController
from user import User


def test_patch_user(filled_controller: CreateUserController) -> None:
    user = {
        "firstName": "John",
        "lastName": "Doe",
        "birth_year": 1990,
        "group": "user"
    }
    assert filled_controller.patch_user(0, user) == True
    assert filled_controller.get_user(0) == User("John", "Doe", 1990, "user", 0)
