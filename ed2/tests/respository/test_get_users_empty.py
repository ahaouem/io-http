from repositories import UserRepository

def test_get_users_empty(empty_repository: UserRepository) -> None:
    assert len(empty_repository.get_users()) == 0
