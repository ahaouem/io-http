class User:
    def __init__(self, firstName: str, lastName: str, birthYear: int, group: str, id: int) -> None:
        self.firstName = firstName
        self.lastName = lastName
        self.birthYear = birthYear
        self.group = group
        self.id = id

    def validate_input(self) -> bool:
        allowed_groups = ["user", "premium", "admin"]
        return self.group in allowed_groups

    def to_dict(self) -> dict:
        return {
            "firstName": self.firstName,
            "lastName": self.lastName,
            "birthYear": self.birthYear,
            "group": self.group,
            "id": self.id
        }

    def __eq__(self, other) -> bool:
        if not isinstance(other, User):
            return False

        return (
            self.firstName == other.firstName and
            self.lastName == other.lastName and
            self.birthYear == other.birthYear and
            self.group == other.group and
            self.id == other.id
        )
