from repositories import UserRepository
from user import User


class CreateUserController:
    def __init__(self, repository: UserRepository) -> None:
        self._repository = repository
        pass

    def get_last_id(self) -> int:
        users = self.get_users()
        id = 0

        try:
            id = users[-1].id
        except:
            pass

        return id

    def validate_id(self, id) -> int:
        users = self.get_users()

        for user in users:
            if user.id == id:
                return True

        return False

    def get_users(self) -> list[User]:
        return self._repository.get_users()

    def get_user(self, id: int) -> User:
        if not self.validate_id(id):
            return None
        return self._repository.get_user(id)

    def post_user(self, user_data: dict) -> bool:
        id = self.get_last_id() + 1
        user = User(user_data['firstName'],
                    user_data["lastName"], user_data["birth_year"], user_data["group"], id)

        if not user.validate_input():
            return False

        return self._repository.add_user(user)

    def patch_user(self, id: int, user_data: dict) -> bool:
        if not self.validate_id(id):
            return False
        return self._repository.patch_user(id, user_data)


    def remove_user(self, id: int) -> bool:
        if not self.validate_id(id):
            return False

        return self._repository.delete_user(id)
