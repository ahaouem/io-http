from controllers import CreateUserControllerw


def test_add_user(empty_controller: CreateUserController) -> None:
    user = {
        "firstName": "John",
        "lastName": "Doe",
        "birth_year": 1990,
        "group": "user"
    }
    assert empty_controller.post_user(user) == True
