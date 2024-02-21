from repositories import UserRepository

def test_get_user(filled_repository: UserRepository) -> None:
    assert filled_repository.get_user(0).lastName == "Doe0"
    assert filled_repository.get_user(4).lastName == "Doe4"
