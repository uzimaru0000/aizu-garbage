-- +migrate Up
CREATE TABLE IF NOT EXISTS place (
  id VARCHAR(6) NOT NULL PRIMARY KEY,
  name VARCHAR(32) NOT NULL
);
INSERT INTO place
values
  ('000100', '相生町（A）');
INSERT INTO place
values
  ('000200', '相生町（B）');
INSERT INTO place
values
  ('000300', '旭町');
INSERT INTO place
values
  ('000400', '居合町');
INSERT INTO place
values
  ('000500', '飯盛1丁目～3丁目');
INSERT INTO place
values
  ('000600', '石堂町');
INSERT INTO place
values
  ('000700', '一箕町（金堀、石ヶ森）');
INSERT INTO place
values
  ('000800', '一箕町（金堀、石ケ森以外）');
INSERT INTO place
values
  ('000900', '上町（A）');
INSERT INTO place
values
  ('001000', '上町（B）');
INSERT INTO place
values
  ('001100', '上町（C）');
INSERT INTO place
values
  ('001200', '駅前町');
INSERT INTO place
values
  ('001300', '追手町');
INSERT INTO place
values
  ('001400', '扇町');
INSERT INTO place
values
  ('001500', '大塚1丁目～2丁目');
INSERT INTO place
values
  ('001600', '大戸町');
INSERT INTO place
values
  ('001700', '大町1丁目～2丁目（A）');
INSERT INTO place
values
  ('001800', '大町1丁目～2丁目（B）');
INSERT INTO place
values
  ('001900', '大町1丁目～2丁目（C）');
INSERT INTO place
values
  ('002000', '御旗町');
INSERT INTO place
values
  ('002100', '表町');
INSERT INTO place
values
  ('005300', '徒之町');
INSERT INTO place
values
  ('005400', '金川町');
INSERT INTO place
values
  ('005500', '川原町');
INSERT INTO place
values
  ('005600', '北青木');
INSERT INTO place
values
  ('005700', '北滝沢1丁目～2丁目');
INSERT INTO place
values
  ('005800', '行仁町（A）');
INSERT INTO place
values
  ('005900', '行仁町（B）');
INSERT INTO place
values
  ('006000', '慶山1丁目～2丁目');
INSERT INTO place
values
  ('006100', '建福寺前');
INSERT INTO place
values
  ('006200', '神指町（南四合）');
INSERT INTO place
values
  ('006300', '神指町（南四合以外）');
INSERT INTO place
values
  ('006400', '高野町');
INSERT INTO place
values
  ('006500', '蚕養町');
INSERT INTO place
values
  ('006600', '材木町1丁目～2丁目');
INSERT INTO place
values
  ('006700', '栄町');
INSERT INTO place
values
  ('006800', '桜町');
INSERT INTO place
values
  ('006900', '五月町');
INSERT INTO place
values
  ('007000', '城東町');
INSERT INTO place
values
  ('007100', '城西町（A）');
INSERT INTO place
values
  ('007200', '城西町（B）');
INSERT INTO place
values
  ('007300', '城南町');
INSERT INTO place
values
  ('007400', '城北町');
INSERT INTO place
values
  ('007500', '昭和町（A）');
INSERT INTO place
values
  ('007600', '昭和町（B）');
INSERT INTO place
values
  ('007700', '城前');
INSERT INTO place
values
  ('007800', '新横町（A）');
INSERT INTO place
values
  ('007900', '新横町（B）');
INSERT INTO place
values
  ('008000', '住吉町');
INSERT INTO place
values
  ('008100', '千石町（A）');
INSERT INTO place
values
  ('008200', '千石町（B）');
INSERT INTO place
values
  ('008300', '宝町');
INSERT INTO place
values
  ('002200', '滝沢町');
INSERT INTO place
values
  ('002300', '館馬町');
INSERT INTO place
values
  ('002400', '館脇町');
INSERT INTO place
values
  ('002500', '中央1丁目（A）');
INSERT INTO place
values
  ('002600', '中央1丁目（B）');
INSERT INTO place
values
  ('002700', '中央1丁目（C）');
INSERT INTO place
values
  ('002800', '中央2丁目（A）');
INSERT INTO place
values
  ('002900', '中央2丁目（B）');
INSERT INTO place
values
  ('003000', '中央3丁目');
INSERT INTO place
values
  ('003001', '対馬舘町');
INSERT INTO place
values
  ('003100', '堤町');
INSERT INTO place
values
  ('003200', '鶴賀町');
INSERT INTO place
values
  ('003300', '天神町（A）');
INSERT INTO place
values
  ('003400', '天神町（B）');
INSERT INTO place
values
  ('003500', '天寧寺町');
INSERT INTO place
values
  ('003600', '中島町');
INSERT INTO place
values
  ('003700', '中町');
INSERT INTO place
values
  ('003800', '七日町');
INSERT INTO place
values
  ('003801', '飯寺北1丁目');
INSERT INTO place
values
  ('003900', '錦町');
INSERT INTO place
values
  ('004000', '西栄町');
INSERT INTO place
values
  ('004100', '西七日町');
INSERT INTO place
values
  ('004200', '西年貢1丁目～2丁目');
INSERT INTO place
values
  ('004300', '日新町');
INSERT INTO place
values
  ('004400', '橋本1丁目～2丁目');
INSERT INTO place
values
  ('004500', '花春町');
INSERT INTO place
values
  ('004600', '花畑東');
INSERT INTO place
values
  ('004700', '花見ケ丘1丁目～3丁目');
INSERT INTO place
values
  ('004800', '馬場町（A）');
INSERT INTO place
values
  ('004900', '馬場町（B）');
INSERT INTO place
values
  ('005000', '馬場本町');
INSERT INTO place
values
  ('005100', '東栄町');
INSERT INTO place
values
  ('005200', '東千石1丁目～3丁目');
INSERT INTO place
values
  ('008400', '東年貢1丁目～2丁目');
INSERT INTO place
values
  ('008500', '東山町（石山天寧）');
INSERT INTO place
values
  ('008600', '東山町（石山天寧以外）');
INSERT INTO place
values
  ('008700', '桧町');
INSERT INTO place
values
  ('008800', '白虎町');
INSERT INTO place
values
  ('008900', '日吉町');
INSERT INTO place
values
  ('009000', '古川町');
INSERT INTO place
values
  ('009100', '平安町');
INSERT INTO place
values
  ('009200', '本町');
INSERT INTO place
values
  ('009300', '幕内東町');
INSERT INTO place
values
  ('009301', '幕内南町');
INSERT INTO place
values
  ('009400', '町北町（石堂・中沢西）');
INSERT INTO place
values
  ('009500', '町北町（藤室）（A）');
INSERT INTO place
values
  ('009600', '町北町（藤室）（B）');
INSERT INTO place
values
  ('009700', '町北町（石堂、中沢西、藤室以外）');
INSERT INTO place
values
  ('009800', '緑町');
INSERT INTO place
values
  ('009900', '湊町');
INSERT INTO place
values
  ('010000', '南千石町（A）');
INSERT INTO place
values
  ('010100', '南千石町（B）');
INSERT INTO place
values
  ('010200', '南花畑');
INSERT INTO place
values
  ('010300', '南町');
INSERT INTO place
values
  ('010400', '宮町');
INSERT INTO place
values
  ('010500', '明和町');
INSERT INTO place
values
  ('010600', '門田町(徳久・飯寺・日吉丑渕)');
INSERT INTO place
values
  ('010700', '門田町（徳久、飯寺、日吉丑渕以外）');
INSERT INTO place
values
  ('010900', '八角町');
INSERT INTO place
values
  ('010800', '柳原町1丁目～4丁目');
INSERT INTO place
values
  ('011000', '山鹿町');
INSERT INTO place
values
  ('011100', '山見町');
INSERT INTO place
values
  ('011200', '湯川町');
INSERT INTO place
values
  ('011300', '湯川南');
INSERT INTO place
values
  ('011400', '八日町');
INSERT INTO place
values
  ('011500', '米代1丁目～2丁目');
INSERT INTO place
values
  ('011600', '和田1丁目～2丁目');
INSERT INTO place
values
  ('011700', '河東地区');
INSERT INTO place
values
  ('011800', '北会津（荒井地区）');
INSERT INTO place
values
  ('011900', '北会津（舘の内・川南地区）');
INSERT INTO place
values
  ('000000', '0');
-- +migrate Down
  DROP TABLE IF EXISTS place;