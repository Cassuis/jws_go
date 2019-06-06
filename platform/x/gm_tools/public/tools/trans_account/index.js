// Generated by CoffeeScript 1.12.7
(function() {
  var AccountKeyInput, AntButton, Api, App, Button, ButtonToolbar, DropdownButton, Icon, Input, ListGroup, ListGroupItem, MenuItem, Modal, ModalTrigger, Nav, NavItem, Navbar, React, ReactBootstrap, Table, Upload, antd;

  antd = require('antd');

  Api = require('../api/api_ajax');

  AccountKeyInput = require('../../common/account_input');

  React = require('react');

  ReactBootstrap = require('react-bootstrap');

  Upload = antd.Upload;

  Icon = antd.Icon;

  AntButton = antd.Button;

  ButtonToolbar = ReactBootstrap.ButtonToolbar;

  Button = ReactBootstrap.Button;

  MenuItem = ReactBootstrap.MenuItem;

  DropdownButton = ReactBootstrap.DropdownButton;

  Table = ReactBootstrap.Table;

  ModalTrigger = ReactBootstrap.ModalTrigger;

  Modal = ReactBootstrap.Modal;

  Navbar = ReactBootstrap.Navbar;

  Nav = ReactBootstrap.Nav;

  NavItem = ReactBootstrap.NavItem;

  Input = ReactBootstrap.Input;

  ListGroup = ReactBootstrap.ListGroup;

  ListGroupItem = ReactBootstrap.ListGroupItem;

  App = React.createClass({displayName: "App",
    getInitialState: function() {
      return {
        select_server: this.props.curr_server,
        player_to_send: "",
        account_data: ""
      };
    },
    isServerRight: function() {
      return (this.state.select_server != null) && (this.state.select_server.serverName != null);
    },
    isAccountRight: function() {
      return (this.refs.accountin != null) && this.refs.accountin.IsRight();
    },
    getLoadingState: function() {
      if (!this.isServerRight() || !this.isAccountRight()) {
        return "disabled";
      }
      return '';
    },
    handleServerChange: function(data) {
      console.log(data);
      this.setState({
        select_server: data
      });
      return setTimeout(this.OnSend, 100);
    },
    handleUserChange: function(data) {
      console.log(data);
      if (data === "") {
        data = "请输入玩家Id";
      }
      return this.setState({
        player_to_send: data
      });
    },
    getServerName: function() {
      if (this.state.select_server != null) {
        return this.state.select_server.name;
      } else {
        return "";
      }
    },
    post: function(url, params) {
      var opt, temp, x;
      temp = document.createElement('form');
      temp.action = url;
      temp.method = 'post';
      temp.style.display = 'none';
      for (x in params) {
        opt = document.createElement('input');
        opt.name = x;
        opt.value = params[x];
        temp.appendChild(opt);
      }
      document.body.appendChild(temp);
      temp.submit();
      return temp;
    },
    query: function() {
      var data;
      data = {
        "params": this.params,
        "key": this.key
      };
      return this.post("../api/v1/command/" + this.state.select_server.serverName + "/" + this.state.player_to_send + "/" + "getAllInfoByAccountID", {
        key: this.props.curr_key,
        params: this.state.send
      });
    },
    mod: function() {
      var api;
      api = new Api();
      console.log(this.state.send);
      return api.Typ("setAllInfoByAccountID").ServerID(this.state.select_server.serverName).AccountID(this.state.player_to_send).Key(this.props.curr_key).Params(this.state.account_data).Do((function(_this) {
        return function(result) {
          console.log("on getAllInfo");
          return console.log(result);
        };
      })(this));
    },
    cleanUnEquipItem: function() {
      var api;
      api = new Api();
      console.log(this.state.send);
      return api.Typ("cleanUnEquipItems").ServerID(this.state.select_server.serverName).AccountID(this.state.player_to_send).Key(this.props.curr_key).Params().Do((function(_this) {
        return function(result) {
          console.log("on getAllInfo");
          return console.log(result);
        };
      })(this));
    },
    handleChangeDataForCommit: function(event) {
      return this.setState({
        account_data: event.target.value
      });
    },
    getDataForCommit: function() {
      return this.state.account_data;
    },
    getServerName: function() {
      var serverName;
      if (this.state.select_server === void 0) {
        serverName = "";
      } else {
        serverName = this.state.select_server.serverName;
      }
      return serverName;
    },
    render: function() {
      return React.createElement("div", null, React.createElement(AccountKeyInput, Object.assign({}, this.props, {
        "ref": "accountin",
        "can_cb": this.handleUserChange
      })), React.createElement(Button, {
        "bsStyle": 'primary',
        "onClick": this.query
      }, "Query"), React.createElement(Upload, Object.assign({}, this.props, {
        "action": "../api/v1/command/" + this.getServerName() + "/" + this.state.player_to_send + "/" + "setAllInfoByAccountID",
        "data": {
          "key": this.props.curr_key,
          "params": []
        }
      }), React.createElement(Button, {
        "type": "ghost"
      }, React.createElement(Icon, {
        "type": "upload"
      }), " 上传账号信息")), React.createElement("hr", null), React.createElement(Button, {
        "bsStyle": 'danger',
        "onClick": this.cleanUnEquipItem
      }, "CleanUnEquipItem"), React.createElement("input", {
        "type": "text",
        "value": this.getDataForCommit(),
        "onChange": this.handleChangeDataForCommit
      }));
    }
  });

  module.exports = App;

}).call(this);