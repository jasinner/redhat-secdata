import xml.etree.ElementTree as ET

NS = '{http://oval.mitre.org/XMLSchema/oval-definitions-5}'

RHCOS_CPES = ['cpe:/o:redhat:enterprise_linux:8::baseos',
    'cpe:/a:redhat:enterprise_linux:8::appstream',
    'cpe:/a:redhat:enterprise_linux:8::nfv',
    'cpe:/a:redhat:rhel_extras_rt:8',
    'cpe:/o:redhat:enterprise_linux:8::fastdatapath',
    'cpe:/a:redhat:openshift:4.4::el8',
    'cpe:/a:redhat:rhosemc:1.0::el8',
    'cpe:/a:redhat:openshift:4.5::el8']

tree = ET.parse('testdata/rhel-8.oval.xml')
root = tree.getroot()
for definition in root.find(NS +'definitions'):
    for meta in definition.find(NS +'metadata'):
        for advisory in meta.iter(NS + 'advisory'):
            for affected_cpe in advisory.iter(NS + 'affected_cpe_list'):
                for cpe in affected_cpe:
                    if cpe.text in RHCOS_CPES:
                        print(definition.attrib)
