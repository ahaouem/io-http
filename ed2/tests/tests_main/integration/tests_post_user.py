import pytest

@pytest.fixture
def client():
    with flask_app.test_client() as testing_client:
        with flask_app.app_context():
            yield testing_client

def test_post_user_integration_invalid_data(client):
    invalid_data = {"firstName": "John", "lastName": "Doe"}  
    response = client.post("/users", json=invalid_data)
    assert response.status_code == BAD_REQUEST

def test_post_user_integration_valid_data(client):
    valid_data = {"firstName": "Jane", "lastName": "Doe", "birth_year": 1992, "group": "B"}
    response = client.post("/users", json=valid_data)
    assert response.status_code == CREATED
