from repositories import UserRepository

def test_user_repository_can_be_instantiated(empty_repository: UserRepository) -> None:
    assert isinstance(empty_repository, UserRepository)
