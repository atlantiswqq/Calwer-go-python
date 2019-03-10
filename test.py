# _*_ coding:utf-8 _*_ 
import time
import requests
import openpyxl
from bs4 import BeautifulSoup
from datetime import date
from datetime import datetime
from datetime import timedelta

class Calwer(object):
    def __init__(self):
        self.home_url = 'https://www.1396j.com/xyft/kaijiang?date=%s&_=%s'
        self.excel_path="/Users/wqq/Desktop/kunge.xlsx"
        
    def get_date(self, start_date="2019-01-01", end_date=None):
        if not end_date:
            end_date = datetime.today().strftime("%Y-%m-%d")
        date_list = []
        datediff = self.string_to_date(end_date)-self.string_to_date(start_date)
        sentence = datediff.days
        while sentence >= 0:
            date_list.append(start_date)
            start_date = self.date_to_string(self.string_to_date(start_date) + timedelta(days=1))
            sentence = (self.string_to_date(end_date) - self.string_to_date(start_date)).days
        return date_list
            
    def date_to_string(self, ddate):
        return ddate.strftime("%Y-%m-%d")

    def string_to_date(self, sdate):
        year_s, mon_s, day_s = sdate.split('-')
        return datetime(int(year_s), int(mon_s), int(day_s))
    
    def get_timestamp(self):
        return int(time.time())
        
    def get_html(self, url):
        response = requests.get(url)
        html = response.content.decode("utf-8")
        return html
    
    def parse_html(self, html):
        soup = BeautifulSoup(html,"lxml")
        tbodys = soup.find_all(name="tbody")
        page_list = []
        for s in tbodys:
            for tr in s.findAll("tr"):
                infors = []
                for td in tr.findAll("td"):
                    print(td.get_text().replace("\xa0"," ").replace("\n"," "))
                    infors.append(td.get_text().replace("\xa0"," ").replace("\n"," "))
                page_list.append(infors)
        return page_list

    def main(self):
        wb = openpyxl.Workbook()
        date_list = self.get_date()
        timestamp = self.get_timestamp()
        for sdate in date_list:
            ws = wb.create_sheet(sdate)
            url = self.home_url % (sdate, timestamp)
            print(url)
            html = self.get_html(url)
            pages = self.parse_html(html)
            for item in pages:
                print(item)
                ws.append(item)
        wb.save(self.excel_path)
        
if __name__ == '__main__':
    c = Calwer()
    c.main()
    