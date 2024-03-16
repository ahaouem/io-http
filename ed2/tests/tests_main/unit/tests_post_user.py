import pytest
from unittest.mock import patch, MagicMock
from main import app as flask_app 
from main import BAD_REQUEST, CREATED

@pytest.fixture
def client():
    with flask_app.test_client() as testing_client:
        with flask_app.app_context():
            yield testing_client

@patch('main.validate_user_data')
@patch('controllers.UserController.post_user')
def test_post_user_invalid_data(mock_post_user, mock_validate_user_data, client):
    mock_validate_user_data.return_value = False
    response = client.post("/users", json={"some": "data"})
    assert response.status_code == BAD_REQUEST
    mock_post_user.assert_not_called()

@patch('main.validate_user_data')
@patch('controllers.UserController.post_user')
def test_post_user_valid_data(mock_post_user, mock_validate_user_data, client):
    mock_validate_user_data.return_value = True
    mock_post_user.return_value = True
    response = client.post("/users", json={"name": "John Doe", "age": 30})
    assert response.status_code == CREATED
    mock_post_user.assert_called_once()

@patch('main.validate_user_data')
@patch('controllers.UserController.post_user')
def test_post_user_creation_failed(mock_post_user, mock_validate_user_data, client):
    mock_validate_user_data.return_value = True
    mock_post_user.return_value = False
    response = client.post("/users", json={"name": "John Doe", "age": 30})
    assert response.status_code == BAD_REQUEST
    mock_post_user.assert_called_once()


