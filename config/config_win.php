<?php
class config_win {
	
	static protected $_root= null;

	static public function root() {
	 if (is_null(self::$_root)) self::$_root= dirname(__FILE__);
	 return self::$_root;
	}

	static public function readConfig(){
		// Code
	}
}