function userClass(username, label) {
    var _name = username;
    var _label = label;
    var _personElement;
    var _chatElement;

    this.name = function () {
        return _name;
    };

    this.label = function () {
        return _label;
    };

    this.personElement = function () {
        return _personElement;
    };

    this.chatElement = function () {
        return _chatElement;
    };

    this.setPersonElement = function (element) {
        _personElement = element;
    };
    
    this.setChatElement = function (element) {
        _chatElement = element;
    };
}