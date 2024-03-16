import pytest

@pytest.fixture
def client():
    with flask_app.test_client() as testing_client:
        with flask_app.app_context():
            yield testing_client

def test_patch_user_integration_invalid_data(client):
    invalid_data = {"firstName": "John"} 
    response = client.patch("/users/1", json=invalid_data)
    assert response.status_code == BAD_REQUEST

def test_patch_user_integration_not_found(client):
    valid_data = {"firstName": "Jane", "lastName": "Doe", "birth_year": 1992, "group": "B"}
    response = client.patch("/users/999", json=valid_data)
    assert response.status_code == NOT_FOUND

def test_patch_user_integration_success(client):
    valid_data = {"firstName": "Jane", "lastName": "Doe", "birth_year": 1992, "group": "B"}
    response = client.patch("/users/1", json=valid_data)
    assert response.status_code == OK
