using DS_Lab1.Models;

namespace DS_Lab1.CommunicationElements
{
    public class PersonRequest
    {
        public string Name { get; set; }
        public int? Age { get; set; }
        public string? Adress { get; set; }
        public string? Work { get; set; }

        public PersonRequest(string name, int? age = null, string? adress = null, string? work = null)
        {
            Name = name;
            Age = age;
            Adress = adress;
            Work = work;
        }

        public Person ToPerson() => new(Name, Age, Adress, Work);
    }
}
