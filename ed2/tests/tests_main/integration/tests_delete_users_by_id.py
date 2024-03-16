import pytest

@pytest.fixture
def client():
    with flask_app.test_client() as testing_client:
        with flask_app.app_context():
            yield testing_client

def test_delete_user_integration_success(client):
    response = client.delete("/users/1")
    assert response.status_code == OK

def test_delete_user_integration_not_found(client):
    response = client.delete("/users/9999")
    assert response.status_code == NOT_FOUND
