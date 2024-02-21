from user import User

users = []


class UserRepository:
    def __init__(self, users: list[User]) -> None:
        self.users: list[User] = users

    def add_user(self, user: User) -> bool:
        self.users.append(user)
        return True

    def get_user(self, id) -> User:
        for user in self.users:
            if user.id == id:
                return user

        return None

    def get_users(self) -> list[User]:
        return self.users

    def patch_user(self, id: int, user_data: dict) -> bool:
        for user in self.users:
            if user.id == id:
                if 'firstName' in user_data:
                    user.firstName = user_data['firstName']
                if 'lastName' in user_data:
                    user.lastName = user_data['lastName']
                if 'birth_year' in user_data:  # Adjust if using 'birthYear' in your User model
                    user.birthYear = user_data['birth_year']
                if 'group' in user_data:
                    user.group = user_data['group']
                return True
        return False  

    def delete_user(self, id: int) -> bool:
        for user in self.users:
            if user.id == id:
                self.users.remove(user)
                return True
        return False

