from controllers import CreateUserController
from user import User


def test_get_user(filled_controller: CreateUserController, filled_repository_users: list[User]) -> None:
    assert filled_controller.get_user(0) == filled_repository_users[0]
    assert filled_controller.get_user(4) == filled_repository_users[4]
