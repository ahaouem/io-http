import pytest

@pytest.fixture
def client():
    with flask_app.test_client() as testing_client:
        with flask_app.app_context():
            yield testing_client

def test_get_user_integration_not_found(client):
    response = client.get("/users/999")
    assert response.status_code == NOT_FOUND

def test_get_user_integration_found(client):
    response = client.get("/users/1")
    assert response.status_code == OK

