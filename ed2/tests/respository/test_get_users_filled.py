from repositories import UserRepository

def test_get_users_filled(filled_repository: UserRepository) -> None:
    assert len(filled_repository.get_users()) == 5
