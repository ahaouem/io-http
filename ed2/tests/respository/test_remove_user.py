from repositories import UserRepository

def test_remove_user(filled_repository: UserRepository) -> None:
    assert filled_repository.delete_user(0) == True
    assert filled_repository.get_user(0) == None
