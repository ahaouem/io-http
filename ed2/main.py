from flask import Flask, request, jsonify, Response
from repositories import UserRepository
from controllers import CreateUserController
from user import User

app = Flask(__name__)

users = []
repository = UserRepository(users)
controller = CreateUserController(repository)

OK = 200
CREATED = 201
BAD_REQUEST = 400
NOT_FOUND = 404

def validate_user_data(user_data: dict) -> bool:
    return "firstName" in user_data and "lastName" in user_data and "birth_year" in user_data and "group" in user_data

@app.get("/users")
def get_users() -> Response:
    users = controller.get_users()
    return jsonify([user.to_dict() for user in users]), OK

@app.get("/users/<int:id>")
def get_user(id: int) -> Response:
    user = controller.get_user(id)
    if user is None:
        return '', NOT_FOUND
    return jsonify(user.to_dict()), OK

@app.post("/users")
def post_user() -> Response:
    body = request.json
    if not validate_user_data(body):
        return '', BAD_REQUEST
    if controller.post_user(body):
        return '', CREATED
    return '', BAD_REQUEST

@app.patch("/users/<int:id>")
def patch_user(id: int) -> Response:
    body = request.json
    if controller.patch_user(id, body):
        return '', OK
    return '', NOT_FOUND

@app.delete("/users/<int:id>")
def delete_user(id: int) -> Response:
    if controller.remove_user(id):
        return '', OK
    return '', NOT_FOUND

if __name__ == "__main__":
    app.run(debug=True)
