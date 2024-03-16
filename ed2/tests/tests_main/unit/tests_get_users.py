import pytest
from unittest.mock import patch
from main import app as flask_app

@pytest.fixture
def client():
    with flask_app.test_client() as testing_client:
        with flask_app.app_context():
            yield testing_client

@patch('controllers.UserController.get_users')
def test_get_users_empty(mock_get_users, client):
    mock_get_users.return_value = []
    response = client.get("/users")
    assert response.status_code == 200
    assert response.json == []

@patch('controllers.UserController.get_users')
def test_get_users_with_data(mock_get_users, client):
    mock_data = [{'id': 1, 'name': 'John Doe'}, {'id': 2, 'name': 'Jane Doe'}]
    mock_get_users.return_value = mock_data
    response = client.get("/users")
    assert response.status_code == 200
    assert response.json == mock_data


