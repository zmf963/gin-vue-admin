import requests
from bs4 import BeautifulSoup
import re
from urllib.parse import quote
import json
import math

try:
    from common.log import logger
except Exception as e:
    print(e)

headers = {'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64; rv:46.0) Gecko/20100101 Firefox/46.0'}

# www.beianbeian.com
def beianbeianApi(domain):
    logger.info('Load beianbeianApi: '+domain)
    # 获取备案ID
    beianId = ''
    url = 'http://www.beianbeian.com/s-0/{}.html'.format(domain)
    try:
        res = requests.get(url=url, headers=headers, allow_redirects=False, verify=False, timeout=10)
    except Exception as e:
        logger.warning('[error] http://www.beianbeian.com is die \n{}'.format(e.args))
        return []

    text = res.text
    soup_1 = BeautifulSoup(text, 'html.parser')
    tbodys = soup_1.find_all('tbody', id='table_tr')
    for tbody in tbodys:
        a_hrefs = tbody.find_all('a')
        for a_href in a_hrefs:
            if '反查' in a_href.get_text():
                beianId = a_href['href']
    if beianId:
        beianSearchUrl = 'http://www.beianbeian.com' + beianId
        logger.info('查询到备案号: {}'.format(beianSearchUrl))
    else:
        logger.info('没有匹配到备案号')
        return []

    # 备案反查域名
    beianbeianNewDomains = []
    tempDict = {}
    # url = r'http://www.beianbeian.com/search-1/%E6%B5%99B2-20080224.html'
    try:
        res = requests.get(url=beianSearchUrl, headers=headers, allow_redirects=False, verify=False, timeout=10)
    except Exception as e:
        logger.warning('[error] request : {}\n{}'.format(beianSearchUrl, e.args))
        return []
    text = res.text
    soup = BeautifulSoup(text, 'html.parser')
    tbodys = soup.find_all('tbody', id='table_tr')
    for tbody in tbodys:
        trs = tbody.find_all('tr')
        for tr in trs:
            tds = tr.find_all('td')
            companyName = tds[4].get_text()
            newDomain = tds[5].get_text().strip().replace('www.', '')
            time = tds[6].get_text()
            if newDomain not in tempDict:
                tempDict[newDomain] = (companyName, newDomain, time)
                beianbeianNewDomains.append((companyName, newDomain, time))
    beianbeianNewDomains = list(set(beianbeianNewDomains))
    logger.info('beianbeianApi去重后共计{}个顶级域名'.format(len(beianbeianNewDomains)))
    return beianbeianNewDomains

# icp.chinaz.com
def chinazApi(domain):

    # 解析chinaz返回结果的json数据
    def parse_json(json_ret):
        chinazNewDomains = []
        results = json_ret['data']
        for result in results:
            companyName = result['webName']
            newDomain = result['host']
            time = result['verifyTime']
            chinazNewDomains.append((companyName, newDomain, time))     # [('城市产业服务平台', 'cloudfoshan.com', '2020-06-09'), ('城市产业服务平台', 'cloudguangzhou.com', '2020-06-09')]
        chinazNewDomains = list(set(chinazNewDomains))
        return chinazNewDomains

    chinazNewDomains = []
    tempDict = {}
    tempList = []

    # 获取域名的公司名字
    url = r'http://icp.chinaz.com/{}'.format(domain)
    try:
        res = requests.get(url=url, headers=headers, allow_redirects=False, verify=False, timeout=10)
    except Exception as e:
        logger.info('[error] request : {}\n{}'.format(url, e.args))
        return [], []
    text = res.text

    companyName = re.search("var kw = '([\S]*)'", text)
    if companyName:
        print(companyName.groups())
        companyName = companyName.group(1)
        logger.info('公司名: {}'.format(companyName))
        companyNameUrlEncode = quote(str(companyName))
    else:
        logger.info('没有匹配到公司名')
        return [], []

    # 备案反查域名
    headers['Content-Type'] = 'application/x-www-form-urlencoded; charset=UTF-8'
    url = 'http://icp.chinaz.com/Home/PageData'
    data = 'pageNo=1&pageSize=20&Kw=' + companyNameUrlEncode
    try:
        res = requests.post(url=url, headers=headers, data=data, allow_redirects=False, verify=False, timeout=10)
    except Exception as e:
        ('[error] request : {}\n{}'.format(url, e.args))
        return [], []

    json_ret = json.loads(res.text)
    if 'amount' not in json_ret.keys():
        return chinazNewDomains, []
    amount = json_ret['amount']
    pages = math.ceil(amount / 20)
    logger.info('页数: {}'.format(pages))
    # 解析返回的json数据包，过滤出公司名，域名，时间     eg: ('城市产业服务平台', 'cloudhuizhou.com', '2020-06-09')
    tempList.extend(parse_json(json_ret))

    # 继续获取后面页数
    for page in range(2, pages+1):
        logger.info('请求第{}页'.format(page))
        data = 'pageNo={}&pageSize=20&Kw='.format(page) + companyNameUrlEncode
        try:
            res = requests.post(url=url, headers=headers, data=data, allow_redirects=False, verify=False, timeout=10)
            json_ret = json.loads(res.text)
            tempList.extend(parse_json(json_ret))
        except Exception as e:
            logger.info('[error] request : {}\n{}'.format(url, e.args))


    for each in tempList:
        if each[1] not in tempDict:
            tempDict[each[1]] = each
            chinazNewDomains.append(each)

    logger.info('chinazApi去重后共计{}个顶级域名'.format(len(chinazNewDomains)))
    return chinazNewDomains, companyName


def run_beian2domain(domain):
    beianNewDomains = []
    # beianbeianNewDomains = beianbeianApi(domain)  # 失效
    chinazNewDomains, companyName = chinazApi(domain)

    tempDict = {}
    for each in chinazNewDomains:
        if each[1] not in tempDict:
            tempDict[each[1]] = each
            beianNewDomains.append(each)

    logger.info('去重后共计{}个顶级域名'.format(len(beianNewDomains)))
    return beianNewDomains

if __name__ == '__main__':
    try:
        import os,sys
        basedir = os.path.dirname(os.path.abspath(__file__))
        sys.path.append(os.path.abspath(os.path.join(basedir, '../../')))
        from common.log import logger
        logger.info("main")
    except Exception as e:
        print(e)
    icp = '京ICP证030173号'
    icp = '京ICP备11041704号'
    logger.info(run_beian2domain(icp))