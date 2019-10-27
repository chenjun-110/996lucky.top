import React from 'react';
import { Form, Icon, Input, Button, message } from 'antd';
import { setToken, getToken } from '../store/index.js'
import api from '../api/index.js'
import {withRouter} from "react-router-dom";
function hasErrors(fieldsError) {
  return Object.keys(fieldsError).some(field => fieldsError[field]);
}

class HorizontalLoginForm extends React.Component {
  constructor(props) {
    super(props);
  }
  componentDidMount() {
    // To disabled submit button at the beginning.
    this.props.form.validateFields();
    console.log('this.props', this.props)
  }

  handleSubmit = e => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        console.log(api)
        api.login({username: values.username, password: values.password}).then((data) => {
           console.log(data)
           if (data.success) {
            message.success(data.message);
            setToken(data.value)
            // this.props.history.push('/users')
           }
        })
      }
    }); 
  };

  render() {
    const { getFieldDecorator, getFieldsError, getFieldError, isFieldTouched } = this.props.form;

    // Only show error after a field is touched.
    const usernameError = isFieldTouched('username') && getFieldError('username');
    const passwordError = isFieldTouched('password') && getFieldError('password');
    return (
      <header className="App-header">
      <Form layout="inline" onSubmit={this.handleSubmit}>
        <Form.Item validateStatus={usernameError ? 'error' : ''} help={usernameError || ''}>
          {getFieldDecorator('username', {
            rules: [{ required: true, message: 'Please input your username!' }],
          })(
            <Input
              prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
              placeholder="Username"
            />,
          )}
        </Form.Item>
        <Form.Item validateStatus={passwordError ? 'error' : ''} help={passwordError || ''}>
          {getFieldDecorator('password', {
            rules: [{ required: true, message: 'Please input your Password!' }],
          })(
            <Input
              prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
              type="password"
              placeholder="Password"
            />,
          )}
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" disabled={hasErrors(getFieldsError())}>
            Log in
          </Button>
        </Form.Item>
      </Form>
      </header>
    );
  }
}

const WrappedHorizontalLoginForm = Form.create({ name: 'horizontal_login' })(HorizontalLoginForm);


export default withRouter(WrappedHorizontalLoginForm);
