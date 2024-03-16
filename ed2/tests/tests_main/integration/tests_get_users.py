import pytest

@pytest.fixture
def client():
    with flask_app.test_client() as testing_client:
        with flask_app.app_context():
            yield testing_client

def test_get_users_integration(client):
    response = client.get("/users")
    assert response.status_code == 200
    assert response.json == []

