import pytest
from unittest.mock import patch
from main import app as flask_app  
from main import BAD_REQUEST, OK, NOT_FOUND

@pytest.fixture
def client():
    with flask_app.test_client() as testing_client:
        with flask_app.app_context():
            yield testing_client

@patch('main.validate_user_data')
@patch('controllers.UserController.patch_user')
def test_patch_user_invalid_data(mock_patch_user, mock_validate_user_data, client):
    mock_validate_user_data.return_value = False
    response = client.patch("/users/1", json={"some": "data"})
    assert response.status_code == BAD_REQUEST
    mock_patch_user.assert_not_called()

@patch('main.validate_user_data')
@patch('controllers.UserController.patch_user')
def test_patch_user_not_found(mock_patch_user, mock_validate_user_data, client):
    mock_validate_user_data.return_value = True
    mock_patch_user.return_value = False
    response = client.patch("/users/1", json={"firstName": "John", "lastName": "Doe", "birth_year": 1990, "group": "A"})
    assert response.status_code == NOT_FOUND

@patch('main.validate_user_data')
@patch('controllers.UserController.patch_user')
def test_patch_user_success(mock_patch_user, mock_validate_user_data, client):
    mock_validate_user_data.return_value = True
    mock_patch_user.return_value = True
    response = client.patch("/users/1", json={"firstName": "Jane", "lastName": "Doe", "birth_year": 1992, "group": "B"})
    assert response.status_code == OK
